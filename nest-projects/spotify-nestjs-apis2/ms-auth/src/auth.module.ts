import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { AuthHelper } from './auth.helper';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
    }),
  ],
  controllers: [AuthController],
  providers: [AuthService, AuthHelper, ConfigService],
})
export class AuthModule {
  constructor(private readonly configService: ConfigService) {
    console.log(
      'ACCESS_TOKEN_SECRET:',
      this.configService.get<string>('ACCESS_TOKEN_SECRET'),
    );
    console.log(
      'ACCESS_TOKEN_EXPIRY:',
      this.configService.get<string>('ACCESS_TOKEN_EXPIRY'),
    );
  }
}
