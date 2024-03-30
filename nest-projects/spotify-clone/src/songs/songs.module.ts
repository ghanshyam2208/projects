import { Module } from '@nestjs/common';
import { SongsController } from './songs.controller';
import { SongsService } from './songs.service';
import { MysqlDummyConnection } from 'src/common/constant/mysql.connection';

@Module({
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
