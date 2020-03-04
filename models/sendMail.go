package models

// ["smtp.163.com", 25,"h1803405349@163.com", "@HFT237854"]

import (
    "fmt"
    "net/smtp"
    "bytes"
    "time"
    "io/ioutil"
    "encoding/base64"
    "strings"
    "log"
)

// define email interface, and implemented auth and send method
type Mail interface {
    Auth()
    Send(message Message) error
}

type SendMail struct {
    user     string
    password string
    host     string
    port     string
    auth     smtp.Auth
    // file_arr []string
}

type Attachment struct {
    name        string
    contentType string
    withFile    bool
}

type Message struct {
    from        string
    from_user   string
    to          []string
    cc          []string
    bcc         []string
    subject     string
    body        string
    contentType string
    attachment  Attachment
}


func SendMailInit(Email_addr string, feedBackMsg string, filePath string) {
    var mail Mail
    // feedBackMsg = mail.feebBackMsg
    // mail = &SendMail{user: "18508105904@163.com", password: "H237854", host: "smtp.163.com", port: "25"}
    // 阿里云屏蔽163的25端口,所以使用qq邮箱
    // mail = &SendMail{user: "18508105904@163.com", password: "H237854", host: "ssl://smtp.163.com", port: "465"}
    mail = &SendMail{user: "1803405349@qq.com", password: "rqlxmwngarcqifhg", host: "smtp.qq.com", port: "587"}
    // fmt.Println('atts=====', atts);
    message := Message{
        from: "1803405349@qq.com",
        // to: []string{"18508105904@163.com"},
        to: []string{"1803405349@qq.com"},
        // to: []string{"mindplus@dfrobot.com"},
        from_user: "Mind+用户反馈",
        cc: []string{},
        bcc: []string{},
        subject: Email_addr+"反馈信息",
        body: `
        <html style="margin:0;padding:0;">
        <body>
        <div style="width:100%;min-height:300px;margin:auto;background:url('http://download3.dfrobot.com.cn/website/image/robot.png') #f8ba0d 95% 90% no-repeat;color:white;padding:1rem 3rem;background-size: 150px;box-sizing: border-box;">
            <h3 style="text-align:center;max-height:50px;position: relative;margin: 2rem 0;">
                <img src="http://download3.dfrobot.com.cn/website/image/logo.png" alt="logo" style="position: absolute;left: 0;top: -25px;" />
                Mind Plus用户反馈信息
            </h3>
            <p><b>发送人</b>：<a href="email">`+Email_addr+`</a></p>
            <p style="width:70%"><b>内容</b>：`+feedBackMsg+`</p>
            <p style="color:#333;font-weight:bold;">请检查附件内是否包含用户反馈文件。</p>
        </div>
        </body>
        </html>
        `,
        contentType: "text/html;charset=utf-8", //内容以html显示
        attachment: Attachment{
            name:        filePath,
            withFile:    true,
        },
    }
    mail.Send(message)
}

func (mail *SendMail) Auth() {
    mail.auth = smtp.PlainAuth("", mail.user, mail.password, mail.host) //获取邮箱验证


}

func (mail SendMail) Send(message Message) error {
    mail.Auth()
    buffer := bytes.NewBuffer(nil)
    boundary := "GoBoundary"
    Header := make(map[string]string)
    Header["From"] = message.from_user+"<"+message.from+">"
    Header["To"] = strings.Join(message.to, ";")
    Header["Cc"] = strings.Join(message.cc, ";")
    Header["Bcc"] = strings.Join(message.bcc, ";")
    Header["Subject"] = message.subject
    Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
    Header["Mime-Version"] = "1.0"
    Header["Date"] = time.Now().String()
    mail.writeHeader(buffer, Header)

    body := "\r\n--" + boundary + "\r\n"
    body += "Content-Type:" + message.contentType + "\r\n"
    body += "\r\n" + message.body + "\r\n"
    buffer.WriteString(body)


    fmt.Println("start send mail2222", smtp.SendMail);
    /*if message.attachment.withFile {
        attachment := "\r\n--" + boundary + "\r\n"
        attachment += "Content-Transfer-Encoding:base64\r\n"
        attachment += "Content-Disposition:attachment\r\n"
        attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + message.attachment.name + "\"\r\n"
        buffer.WriteString(attachment)
        defer func() {
            if err := recover(); err != nil {
                log.Fatalln(err)
            }
        }()

        file_arr := strings.Split(message.attachment.name, ",")
        for key, value := range file_arr {
            fmt.Println(file_arr[key])
            mail.writeFile(buffer, value)
        }
    }
*/
    // 多附件
    if message.attachment.withFile {
        attachment := "\r\n--" + boundary + "\r\n"
        file_arr := strings.Split(message.attachment.name, ",")
        attachment += "Content-Transfer-Encoding:base64\r\n"
        attachment += "Content-Disposition:attachment\r\n"
        for key, value := range file_arr {
            attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + value + "\"\r\n"
            buffer.WriteString(attachment)
            defer func() {
                if err := recover(); err != nil {
                    log.Fatalln(err)
                }
            }()
            fmt.Println(file_arr[key])
            mail.writeFile(buffer, value)
        }
    }

    buffer.WriteString("\r\n--" + boundary + "--")
    smtp.SendMail(mail.host+":"+mail.port, mail.auth, message.from, message.to, buffer.Bytes())
    fmt.Println("***send mail success")
    return nil
}

func (mail SendMail) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
    header := ""
    for key, value := range Header {
        header += key + ":" + value + "\r\n"
    }
    header += "\r\n"
    buffer.WriteString(header)
    return header
}

// read and write the file to buffer
func (mail SendMail) writeFile(buffer *bytes.Buffer, fileName string) {
    file, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err.Error())
    }
    payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
    base64.StdEncoding.Encode(payload, file)
    buffer.WriteString("\r\n")
    for index, line := 0, len(payload); index < line; index++ {
        buffer.WriteByte(payload[index])
        if (index+1)%76 == 0 {
            buffer.WriteString("\r\n")
        }
    }
}
