package tests

import (
	"testing"

	"../scan"
)

func TestExtractLink(t *testing.T) {

	testStr := `t: 100%;color: #656565;font-weight: normal;text-decoration: underline;">up=date your preferences
	</a> or <a href=3D"https://crumbsandwhiskers.us9.list=-manage.com/unsubscribe?u=3D06fdf6bf53ccd492af7d57869&id=3Dcf640c7284&e=3D=a44704fa8e&c=3D429abddeb5" style=3D"mso-line-height-rule: exactly;-ms-text-s=ize-adjust: 100%;
	-webkit-text-size-adjust: 100%;color: #656565;font-weight=`

	expectedResult := `"https://crumbsandwhiskers.us9.list=-manage.com/unsubscribe?u=3D06fdf6bf53ccd492af7d57869&id=3Dcf640c7284&e=3D=a44704fa8e&c=3D429abddeb5"`

	actualResult := scan.ExtractLink(testStr)

	if actualResult != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}

}

func TestExtractLinks(t *testing.T) {

	testStr := `t: 100%;color: #656565;font-weight: normal;text-decoration: underline;">up=date your preferences
	</a> or <a href=3D"https://crumbslink0.us9.list=-manage.com/unsubscribe?u=3D06fdf6bf53ccd492af7d57869&id=3Dcf640c7284&e=3D=a44704fa8e&c=3D429abddeb5" style=3D"mso-line-height-rule: exactly;-ms-text-s=ize-adjust: 100%;
	-webkit-text-size-adjust: 100%;color: #656565;font-weight=t: 100%;color: #656565;font-weight: normal;text-decoration: underline;">up=date your preferences
	</a> or <a href=3D"https://crumbslink1.us9.list=-manage.com/unsubscribe?u=3D06fdf6bf53ccd492af7d57869&id=3Dcf640c7284&e=3D=a44704fa8e&c=3D429abddeb5" style=3D"mso-line-height-rule: exactly;-ms-text-s=ize-adjust: 100%;
	-webkit-text-size-adjust: 100%;color: #656565;font-weight=t: 100%;color: #656565;font-weight: normal;text-decoration: underline;">up=date your preferences
	</a> or <a href=3D"https://crumbslink2.us9.list=-manage.com/unsubscribe?u=3D06fdf6bf53ccd492af7d57869&id=3Dcf640c7284&e=3D=a44704fa8e&c=3D429abddeb5" style=3D"mso-line-height-rule: exactly;-ms-text-s=ize-adjust: 100%;
	-webkit-text-size-adjust: 100%;color: #656565;font-weight=`

	expectedResult := [3]string{
		`"https://crumbslink0.us9.list=-manage.com/unsubscribe?u=3D06fdf6bf53ccd492af7d57869&id=3Dcf640c7284&e=3D=a44704fa8e&c=3D429abddeb5"`,
		`"https://crumbslink1.us9.list=-manage.com/unsubscribe?u=3D06fdf6bf53ccd492af7d57869&id=3Dcf640c7284&e=3D=a44704fa8e&c=3D429abddeb5"`,
		`"https://crumbslink2.us9.list=-manage.com/unsubscribe?u=3D06fdf6bf53ccd492af7d57869&id=3Dcf640c7284&e=3D=a44704fa8e&c=3D429abddeb5"`,
	}

	actualResult := scan.ExtractAllLinks(testStr)

	for i, _ := range expectedResult {
		if actualResult[i] != expectedResult[i] {
			t.Errorf("Expected %s but got %s", expectedResult, actualResult)
		}
	}

}

func TestExtractAllUnsubLinks(t *testing.T) {

	testStr := `<a href="http://click.email.anthem.com/?qs=16cb" style="color:#37475A;" target="_blank">about us</a>
		<a   href="http://click.email.anthem.com/?qs=16cb" style="color:#37475A;" target="_blank">anthem.com/ca.</a>
		<a   href="http://click.email.anthem.com/?qs=16cb" style="color:#37475A;" target="_blank">Unsubscribe.</a>
		<a href=3D"http://twitter.com/bandmix">http://twitter.com/bandmix</a>
		<a href='http://www.eventbrite.com/org/8184527438' style='color: #666;'>Relish Dating</a>
		<a style='color: #666;' href='http://www.eventbrite.com/inviteunsubscribe?email=jaymo3141%40gmail.com'>unsubscribe</a>`

	expectedResult := [2]string{
		`<a   href="http://click.email.anthem.com/?qs=16cb" style="color:#37475A;" target="_blank">Unsubscribe.</a>`,
		`<a style='color: #666;' href='http://www.eventbrite.com/inviteunsubscribe?email=jaymo3141%40gmail.com'>unsubscribe</a>`,
	}

	actualResult := scan.ExtractAllUnsubLinks(testStr)

	for i, _ := range expectedResult {
		if actualResult[i] != expectedResult[i] {
			t.Errorf("Expected %s but got %s", expectedResult, actualResult)
		}
	}

}

func Test_ExtractUrl_Single_Quote(t *testing.T) {

	singleQuote := `<a href='http://click.email.anthem.com/?qs=16cb' style="color:#37475A;" target="_blank">about us</a>`

	expectedResult := `http://click.email.anthem.com/?qs=16cb`

	actualResult, _ := scan.ExtractUrl(singleQuote)

	if actualResult != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}

}

func Test_ExtractUrl_Double_Quote(t *testing.T) {

	dblQuote := `<a href="http://click.email.anthem.com/?qs=16cb" style="color:#37475A;" target="_blank">about us</a>`

	expectedResult := `http://click.email.anthem.com/?qs=16cb`

	actualResult, _ := scan.ExtractUrl(dblQuote)

	if actualResult != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, actualResult)
	}

}
