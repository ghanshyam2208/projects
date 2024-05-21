import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ClientsModule, Transport } from '@nestjs/microservices';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'ANALYTICS',
        transport: Transport.TCP,
        options: {
          host: '0.0.0.0',
          port: 9003,
        },
      },
      {
        name: 'COMMUNICATION',
        transport: Transport.TCP,
        options: {
          host: '0.0.0.0',
          port: 9002,
        },
      },
    ]),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
