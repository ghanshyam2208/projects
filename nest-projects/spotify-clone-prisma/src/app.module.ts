import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { CommonModule } from './common/common.module';
import { SongsModule } from './songs/songs.module';

@Module({
  imports: [CommonModule, SongsModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
