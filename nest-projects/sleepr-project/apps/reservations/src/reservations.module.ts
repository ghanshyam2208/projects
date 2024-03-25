import { Module } from '@nestjs/common';
import { ReservationsController } from './reservations.controller';
import { ReservationsService } from './reservations.service';
import { DatabaseModule } from '@app/common';
import {
  ReservationsDocument,
  ReservationsSchema,
} from './model/reservations.model';
import { ReservationsRepository } from './reservations.repository';
import { LoggerModule } from '@app/common';

@Module({
  imports: [
    DatabaseModule,
    DatabaseModule.forFeature([
      {
        name: ReservationsDocument.name,
        schema: ReservationsSchema,
      },
    ]),
    LoggerModule,
  ],
  controllers: [ReservationsController],
  providers: [ReservationsService, ReservationsRepository],
})
export class ReservationsModule {}
