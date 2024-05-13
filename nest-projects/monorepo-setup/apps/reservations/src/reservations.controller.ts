import { Body, Controller, Post } from '@nestjs/common';
import { ReservationsService } from './reservations.service';
import {
  CreateReservationPayload,
  ReservationValidationPipe,
} from './reservations.validations';

@Controller('reservations')
export class ReservationsController {
  constructor(private readonly reservationsService: ReservationsService) {}

  @Post()
  createReservation(
    @Body(new ReservationValidationPipe())
    createReservationPayload: CreateReservationPayload,
  ) {
    return this.reservationsService.createReservation(createReservationPayload);
  }
}
