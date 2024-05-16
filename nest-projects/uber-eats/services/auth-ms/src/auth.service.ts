import { Injectable } from '@nestjs/common';
import { CustomConfigService } from './custom-config/custom-config.service';

@Injectable()
export class AuthService {
  constructor(private readonly customConfigService: CustomConfigService) {}

  getHello() {
    return 'hello from auth';
  }
}
