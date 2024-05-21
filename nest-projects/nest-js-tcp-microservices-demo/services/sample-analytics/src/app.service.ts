import { Injectable } from '@nestjs/common';
import { CreateUserEvent } from './user.validation';

@Injectable()
export class AppService {
  private readonly analyticsData: any[] = [];

  handleUserCreated(createUserEvent: CreateUserEvent) {
    console.log('userCreated received');
    this.analyticsData.push({
      ...createUserEvent,
      timestamp: new Date(),
    });
  }

  getAnalytics() {
    return this.analyticsData;
  }
}
