import { Injectable } from '@nestjs/common';
import * as nodemailer from 'nodemailer';

export interface UserCreatedOtpEvent {
  userId: string;
  email: string;
  emailVerificationOtp: string;
}

@Injectable()
export class NotificationHelper {
  private transporter: nodemailer.Transporter;

  constructor() {}

  async sendMail(from: string, to: string, subject: string, otp: string) {
    this.transporter = nodemailer.createTransport({
      host: '0.0.0.0',
      port: 1025,
    });
    // Generate some random content for the email
    const randomContent = `
  <p>Thank you for signing up!</p>
  <p>Your OTP is: <strong>${otp}</strong></p>
  <p>Please use this OTP to verify your email address.</p>
`;

    const mailOptions = {
      from,
      to,
      subject,
      html: `
    <html>
      <head>
        <style>
          body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
          }
          .container {
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            background-color: #fff;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
          }
          h1 {
            color: #333;
          }
          p {
            color: #666;
          }
          strong {
            color: #ff6f61;
          }
        </style>
      </head>
      <body>
        <div class="container">
          <h1>${subject}</h1>
          ${randomContent}
        </div>
      </body>
    </html>
  `,
    };

    try {
      const info = await this.transporter.sendMail(mailOptions);
      console.log(`Email sent: ${info.response}`);
    } catch (error) {
      console.error(error);
    }
  }
}
