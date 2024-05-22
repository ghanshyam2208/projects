import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { CustomConfigModule } from './custom-config/custom-config.module';

@Module({
  imports: [CustomConfigModule],
  controllers: [AuthController],
  providers: [AuthService],
})
export class AuthModule {}
