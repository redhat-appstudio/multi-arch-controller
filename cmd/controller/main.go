package main

import (
	"flag"
	mpcmetrics "github.com/konflux-ci/multi-platform-controller/pkg/metrics"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/metrics"

	zap2 "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"k8s.io/klog/v2"

	// needed for hack/update-codegen.sh
	_ "k8s.io/code-generator"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/konflux-ci/multi-platform-controller/pkg/controller"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	//+kubebuilder:scaffold:imports
	"github.com/go-logr/logr"
)

var (
	mainLog logr.Logger
)

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	var abAPIExportName string
	var logLevel string
	var stackTraceLevel string
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.StringVar(&abAPIExportName, "api-export-name", "jvm-build-service", "The name of the jvm-build-service APIExport.")

	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")

	// logging vars
	flag.StringVar(&logLevel, "zap-log-level", "", "Zap Level to configure the verbosity of logging")
	flag.StringVar(&stackTraceLevel, "zap-stacktrace-level", "", "Zap Level at and above which stacktraces are captured")

	zapFlagSet := flag.NewFlagSet("zap", flag.ContinueOnError)
	opts := zap.Options{
		TimeEncoder: zapcore.RFC3339TimeEncoder,
		ZapOpts:     []zap2.Option{zap2.WithCaller(true)},
	}
	opts.BindFlags(zapFlagSet)
	klog.InitFlags(zapFlagSet)

	flag.Parse()

	setFlagIfNotEmptyOrPanic(zapFlagSet, "zap-log-level", logLevel)
	setFlagIfNotEmptyOrPanic(zapFlagSet, "zap-stacktrace-level", stackTraceLevel)

	logger := zap.New(zap.UseFlagOptions(&opts))
	ctrl.SetLogger(logger)
	klog.SetLoggerWithOptions(logger, klog.ContextualLogger(true))

	ctx := ctrl.SetupSignalHandler()
	restConfig := ctrl.GetConfigOrDie()

	var mgr ctrl.Manager
	var err error
	mopts := ctrl.Options{
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "5483be8f.redhat.com",
		Metrics:                metricsserver.Options{BindAddress: metricsAddr},
	}

	mainLog = ctrl.Log.WithName("main")
	mainLog.Info("creating standard manager")
	mgr, err = controller.NewManager(restConfig, mopts)
	if err != nil {
		mainLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = mpcmetrics.RegisterCommonMetrics(ctx, metrics.Registry); err != nil {
		mainLog.Error(err, "failed to register common metrics")
		os.Exit(1)
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		mainLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		mainLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	mainLog.Info("starting manager")
	if err := mgr.Start(ctx); err != nil {
		mainLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func setFlagIfNotEmptyOrPanic(fs *flag.FlagSet, name, value string) {
	if len(value) > 0 {
		err := fs.Set(name, value)
		if err != nil {
			panic(err)
		}
	}
}
