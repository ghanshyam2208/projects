import { Module } from '@nestjs/common';
import { SongsController } from './songs.controller';
import { SongsService } from './songs.service';
import { MysqlDummyConnection } from 'src/common/constant/mysql.connection';
import { TypeOrmModule } from '@nestjs/typeorm';
import { SongsModel } from './songs.model';

@Module({
  imports: [TypeOrmModule.forFeature([SongsModel])],
  controllers: [SongsController],
  providers: [
    {
      provide: SongsService,
      useClass: SongsService,
    },
    {
      provide: 'MYSQL_CONNECTION',
      useValue: MysqlDummyConnection, // value provider to provide, to be injected in songs service
    },
  ],
})
export class SongsModule {}
