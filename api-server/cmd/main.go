package main

import (
	apiserver "github.com/ajatprabha/operator-example/api-server"
	deploymentsv1alpha1 "github.com/ajatprabha/operator-example/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	kubelog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

var (
	runLog = kubelog.Log.WithName("darkroom-cp").WithName("run")
	scheme = runtime.NewScheme()
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	_ = deploymentsv1alpha1.AddToScheme(scheme)
}

func main() {
	kubelog.SetLogger(zap.New(zap.UseDevMode(true)))

	mgr, err := apiserver.NewManager(ctrl.GetConfigOrDie(), apiserver.Options{
		Scheme:         scheme,
		Port:           5000,
		AllowedDomains: []string{},
	})
	if err != nil {
		runLog.Error(err, "unable to create api-server manager")
		os.Exit(1)
	}

	runLog.Info("starting api-server manager")
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		runLog.Error(err, "problem running api-server manager")
		os.Exit(1)
	}
}
