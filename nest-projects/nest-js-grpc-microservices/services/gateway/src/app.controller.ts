import { Body, Controller, Get, Post } from '@nestjs/common';
import { AppService } from './app.service';
import { CreateUserPayload } from 'proto/auth';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Post()
  createUser(@Body() createUserPayload: CreateUserPayload) {
    return this.appService.createUser(createUserPayload);
  }

  @Get()
  fundAllUsers() {
    return this.appService.fundAllUsers();
  }
}
