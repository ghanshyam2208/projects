import { Controller } from '@nestjs/common';
import { EventPattern } from '@nestjs/microservices';
import { UserCreatedOtpEvent } from './notification.helper';
import { NotificationService } from './notification.service';

@Controller('notification')
export class NotificationController {
  constructor(private readonly notificationService: NotificationService) {}
  @EventPattern('user_created_otp')
  handleUserCreated(userCreatedOtpEvent: UserCreatedOtpEvent) {
    console.log('event received', userCreatedOtpEvent);
    this.notificationService.handleUserCreated(userCreatedOtpEvent);
  }
}
