import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import {
  USERS_SERVICE_NAME,
  UsersServiceClient,
  CreateUserPayload,
} from 'proto/auth';

@Injectable()
export class AppService implements OnModuleInit {
  private userService: UsersServiceClient;

  constructor(@Inject('AUTH_SERVICE') private clientGrpc: ClientGrpc) {}

  onModuleInit() {
    this.userService =
      this.clientGrpc.getService<UsersServiceClient>(USERS_SERVICE_NAME);
  }

  createUser(createUserPayload: CreateUserPayload) {
    console.log('calling in gateway ms user');
    return this.userService.createUser(createUserPayload);
  }

  fundAllUsers() {
    return this.userService.findAllUsers({});
  }
}
