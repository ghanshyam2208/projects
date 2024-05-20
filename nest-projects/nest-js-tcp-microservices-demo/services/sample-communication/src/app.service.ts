import { Injectable } from '@nestjs/common';
import { CreateUserEvent } from './user.validation';

@Injectable()
export class AppService {
  getHello(): string {
    return 'Hello World!';
  }

  handleUserCreated(createUserEvent: CreateUserEvent) {
    console.log(`communication service: ${JSON.stringify(createUserEvent)}`);
  }
}
