import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ComputerModule } from './computer/computer.module';
import { DiskModule } from './disk/disk.module';
import { CpuModule } from './cpu/cpu.module';
import { PowerModule } from './power/power.module';

@Module({
  imports: [ComputerModule, DiskModule, CpuModule, PowerModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
