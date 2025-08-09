package entity

import (
	"html/template"
	"log"
)

type PhishingEmailParams struct {
	EmployeeName   string
	CompanyName    string
	DownloadLink   string
	SenderAddress  string
	Deadline       string
	AttachmentName string
}

const EmailTemplate = `
Hi {{.EmployeeName}},

Our records indicate that your Q3 expense report contains discrepancies that require correction. 
Please review the attached file and submit updates by {{.Deadline}} to avoid payroll delays.

📎 Download your report here: {{.DownloadLink}} (filename: {{.AttachmentName}})

Note: If you did not submit expenses this quarter, please forward this email to IT for investigation.

Best regards,
Corporate Accounting
{{.CompanyName}}`

func GetInvoiceAttack() AttackType {
	t, err := template.New("invoice").Parse(EmailTemplate)
	if err != nil {
		log.Print(err)
	}

	return AttackType{
		ID: "123",
		Params: map[string]any{
			"subject": "Invoice",
		},
		Body: t,
	}
}
