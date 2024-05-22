import { Module } from '@nestjs/common';
import { UserManagementController } from './userManagement.controller';
import { UserManagementService } from './userManagement.service';
import { CustomConfigModule } from './custom-config/custom-config.module';

@Module({
  imports: [CustomConfigModule],
  controllers: [UserManagementController],
  providers: [UserManagementService],
})
export class AppModule {}
