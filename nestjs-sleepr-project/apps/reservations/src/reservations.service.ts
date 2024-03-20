import { Injectable } from '@nestjs/common';
import { CreateReservationDto } from './dto/CreateReservationDto';
import { ReservationsRepository } from './reservations.repository';

@Injectable()
export class ReservationsService {
  constructor(
    private readonly reservationsRepository: ReservationsRepository,
  ) {}
  createReservation(createReservationDto: CreateReservationDto) {
    return this.reservationsRepository.create({
      ...createReservationDto,
      userId: '123',
      timestamp: new Date(),
    });
  }
}
