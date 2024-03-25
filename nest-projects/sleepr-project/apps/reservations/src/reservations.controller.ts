import {
  Body,
  Controller,
  Get,
  Param,
  Post,
  Patch,
  Delete,
} from '@nestjs/common';
import { ReservationsService } from './reservations.service';
import { ReservationsDocument } from './model/reservations.model';
import { CreateReservationDto } from './dto/CreateReservationDto';
import { UpdateReservationDto } from './dto/UpdateReservationDto';
import { ReservationValidationPipe } from './reservation.validation';

@Controller('reservations')
export class ReservationsController {
  constructor(private readonly reservationsService: ReservationsService) {}

  @Get('')
  findAll(): Promise<ReservationsDocument[]> {
    return this.reservationsService.findAll();
  }

  @Post('')
  createReservation(
    @Body(new ReservationValidationPipe())
    createReservationDto: CreateReservationDto,
  ): Promise<ReservationsDocument> {
    return this.reservationsService.createReservation(createReservationDto);
  }

  @Get(':id')
  findOne(@Param('id') _id: string) {
    return this.reservationsService.findOne(_id);
  }

  @Patch(':id')
  update(
    @Param('id') id: string,
    @Body() updateReservationDto: UpdateReservationDto,
  ) {
    return this.reservationsService.update(id, updateReservationDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.reservationsService.remove(id);
  }
}
