package verifycode

import (
	"bytes"
	"fmt"
	"html/template"
)

func newHTMLEmail(code string) (subject string, body string, err error) {
	mailData := &struct {
		Code string
	}{
		Code: code,
	}

	subject = fmt.Sprint("adminx 验证码")

	body, err = getEmailHTMLContent(mailTemplate, mailData)
	return
}

func getEmailHTMLContent(mailTpl string, mailData interface{}) (string, error) {
	tpl, err := template.New("email tpl").Parse(mailTpl)
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	err = tpl.Execute(buffer, mailData)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

const mailTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<base target="_blank" />
		<style type="text/css">
			::-webkit-scrollbar{ display: none; }
		</style>
		<style	style id="cloudAttachStyle" type="text/css">
			#divNeteaseBigAttach, #divNeteaseBigAttach_bak{display:none;}
		</style>
		<style id="blockquoteStyle" type="text/css">blockquote{display:none;}</style>
	</head>
	<body tabindex="0" role="listitem">
		<div id="content" class="netease_mail_readhtml netease_mail_readhtml_webmail">
			<table style="background-color: #e9e9e9; line-height: 1.6; padding-top: 50px; padding-bottom: 50px; color: #666666; font-size: 13px; font-family: Helvetica Neue,Helvetica,Arial,sans-serif;" width="100%" border="0" cellspacing="0" cellpadding="0">
				<tbody>
					<tr>
						<td>
							<table style="padding-left: 30px; padding-right: 30px; padding-top: 27px; padding-bottom: 27px; background-color: #ffffff;" width="700px" border="0" cellspacing="0" cellpadding="0" align="center">
								<tbody>
									<tr>
										<td style="border-bottom: 1px solid #D8D8D8; padding-bottom: 20px; padding-left: 20px;">
											<img style="width: 104px; height:auto" src="https://www.agora.io/en/wp-content/themes/agora-2020/images/agora-logo.svg" alt="">
										</td>
										<td style="width: 200px; border-bottom: 1px solid #D8D8D8; padding-bottom: 20px; text-align: right; padding-right: 20px;">
											Enable Real-time Internet
										</td>
									</tr>
									<tr>
										<td style="padding-left: 20px; padding-right: 20px; padding-top: 50px; padding-bottom: 50px;" colspan="2">
											<table width="100%" border="0" cellspacing="0" cellpadding="0">
												<tbody>
													<tr>
														<td style="padding-bottom: 22px;">
															Dear Customer,
														</td>
													</tr>
													<tr>
														<td>
															Your Agora verification code is {{.Code}}. This code will expire in 10 minutes.
															<br>
															If you didn't ask for this verification code, please ignore this notification. 

														</td>
													</tr>
													<tr>
														<td style="padding-top: 22px;">
															adminx
														</td>
													</tr>	
												</tbody>
											</table>
										</td>
									</tr>
									<tr>
										<td style="padding-bottom: 20px;" colspan="2">
											<table style="border-bottom: 1px solid #D8D8D8; border-top: 1px solid #D8D8D8; padding-bottom: 15px; padding-top: 15px; padding-left: 22px; padding-right: 22px;" width="100%" border="0" cellspacing="0" cellpadding="0">
												<tbody>
													<tr>
														<td>
															<a style="color: #666666; text-decoration: none;" title="Adminx" href="//bytego.dev/">AdminX</a>
														</td>
														<td style="text-align: right;">
															<a style="color: #666666; text-decoration: none; margin-left: 5px;" title="AdminX Console" href="//bytego.dev">Console</a>
															&nbsp;&nbsp;&nbsp;/ 
															<a style="color: #666666; text-decoration: none; margin-left: 5px;" title="Developer Center" href="//bytego.dev">Developer Center</a>
															&nbsp;&nbsp;&nbsp;/ 
															<a style="color: #666666; text-decoration: none; margin-left: 5px;" title="AdminX Community" href="//bytego.dev">AdminX Community</a>
														</td>
													</tr>
												</tbody>
											</table>
										</td>
									</tr>
								</tbody>
							</table>
						</td>
					</tr>
				</tbody>
			</table>
			<div style="clear:both;height:1px;"></div>
		</div>
		<script>
			var _c=document.getElementById('content');
			_c.innerHTML=(_c.innerHTML||'').replace(/(href|formAction|onclick|javascript)/ig, '__$1').replace(/<\/?marquee>/ig,'');
			var _s = _c.getElementsByTagName('style');
			for(var i=0;i<_s.length;i++){ var _st = _s[i].innerHTML.split('}'); for(var j=0;j<_st.length-1;j++){ _st[j] = '.netease_mail_readhtml '+_st[j]; } _s[i].innerHTML = _st.join('}'); }
		</script>
		<style type="text/css">
			* {
  				white-space: normal !important;
  				word-break: break-word !important;
			}
			body{font-size:14px;font-family:arial,verdana,sans-serif;line-height:1.666;padding:0;margin:0;overflow:auto;white-space:normal;word-wrap:break-word;min-height:100px}
			td, input, button, select, body{font-family:Helvetica, 'Microsoft Yahei', verdana}
			pre {white-space:pre-wrap !important;white-space:-moz-pre-wrap;white-space:-pre-wrap;white-space:-o-pre-wrap;word-wrap:break-word;width:95%}
			th,td{font-family:arial,verdana,sans-serif;line-height:1.666}
			img{ border:0}
			header,footer,section,aside,article,nav,hgroup,figure,figcaption{display:block}
			blockquote{margin-right:0px}
		</style>
		<style id="ntes_link_color" type="text/css">a,td a{color:#3370FF}</style>
	</body>
</html>
`
