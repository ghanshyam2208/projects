import { Controller } from '@nestjs/common';
import { EventPattern } from '@nestjs/microservices';

@Controller('notification')
export class NotificationController {
  @EventPattern('login')
  handleUserCreated(createUserEvent: any) {
    console.log('event received', createUserEvent);
  }
}
