import { Inject, Injectable } from '@nestjs/common';
import { CreateUserEvent, UserCreatePayload } from './user.validation';
import { ClientProxy } from '@nestjs/microservices';

@Injectable()
export class AppService {
  private readonly users: UserCreatePayload[] = [];

  constructor(
    @Inject('COMMUNICATION') private readonly communicationClient: ClientProxy,
    @Inject('ANALYTICS') private readonly analyticsClient: ClientProxy,
  ) {}

  getHello(): string {
    return 'Hello World!';
  }

  createUser(userCreatePayload: UserCreatePayload) {
    this.users.push(userCreatePayload);
    console.log('sending email: ', userCreatePayload.email);
    this.communicationClient.emit(
      'user_created',
      new CreateUserEvent(userCreatePayload.email),
    );
    this.analyticsClient.emit(
      'user_created',
      new CreateUserEvent(userCreatePayload.email),
    );
    return userCreatePayload;
  }

  getAnalytics(): any {
    return this.analyticsClient.send({ cmd: 'get_analytics' }, {});
  }
}
