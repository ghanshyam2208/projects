import {
  Body,
  Controller,
  Get,
  Post,
  Patch,
  Param,
  Delete,
} from '@nestjs/common';
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

  @Get()
  getReservations() {
    return this.reservationsService.getReservations();
  }

  @Get(':id')
  getReservationById(@Param('id') id: string) {
    return this.reservationsService.getReservationById(id);
  }

  @Patch(':id')
  updateReservationById(
    @Param('id') id: string,
    @Body(new ReservationValidationPipe())
    createReservationPayload: CreateReservationPayload,
  ) {
    return this.reservationsService.updateReservation(
      id,
      createReservationPayload,
    );
  }

  @Delete(':id')
  deleteReservationById(@Param('id') id: string) {
    return this.reservationsService.deleteReservationById(id);
  }
}
