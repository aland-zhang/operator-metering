package framework

import (
	"time"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	chargebackv1alpha1 "github.com/operator-framework/operator-metering/pkg/apis/chargeback/v1alpha1"
)

func (f *Framework) CreateChargebackReport(report *chargebackv1alpha1.Report) error {
	_, err := f.ChargebackClient.Reports(f.Namespace).Create(report)
	return err
}

func (f *Framework) CreateChargebackScheduledReport(report *chargebackv1alpha1.ScheduledReport) error {
	_, err := f.ChargebackClient.ScheduledReports(f.Namespace).Create(report)
	return err
}

func (f *Framework) GetChargebackScheduledReport(name string) (*chargebackv1alpha1.ScheduledReport, error) {
	return f.ChargebackClient.ScheduledReports(f.Namespace).Get(name, meta.GetOptions{})
}

func (f *Framework) NewSimpleReport(name, queryName string, start, end time.Time) *chargebackv1alpha1.Report {
	return &chargebackv1alpha1.Report{
		ObjectMeta: meta.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Spec: chargebackv1alpha1.ReportSpec{
			ReportingStart:      meta.Time{start},
			ReportingEnd:        meta.Time{end},
			GenerationQueryName: queryName,
			RunImmediately:      true,
		},
	}
}

func (f *Framework) NewSimpleScheduledReport(name, queryName string, lastReportTime time.Time) *chargebackv1alpha1.ScheduledReport {
	return &chargebackv1alpha1.ScheduledReport{
		ObjectMeta: meta.ObjectMeta{
			Name:      name,
			Namespace: f.Namespace,
		},
		Spec: chargebackv1alpha1.ScheduledReportSpec{
			GenerationQueryName: queryName,
			Schedule: chargebackv1alpha1.ScheduledReportSchedule{
				Period: chargebackv1alpha1.ScheduledReportPeriodHourly,
			},
		},
		Status: chargebackv1alpha1.ScheduledReportStatus{
			LastReportTime: &meta.Time{lastReportTime},
		},
	}
}
