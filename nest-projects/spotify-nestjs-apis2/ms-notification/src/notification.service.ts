import { Injectable } from '@nestjs/common';
import { NotificationHelper, UserCreatedOtpEvent } from './notification.helper';

@Injectable()
export class NotificationService {
  constructor(private readonly notificationHelper: NotificationHelper) {}

  async handleUserCreated(userCreatedOtpEvent: UserCreatedOtpEvent) {
    // Send a notification to the user
    await this.notificationHelper.sendMail(
      'server@spotify.com',
      userCreatedOtpEvent.email,
      'Welcome to Spotify',
      userCreatedOtpEvent.emailVerificationOtp,
      userCreatedOtpEvent.userId,
    );
  }
}
