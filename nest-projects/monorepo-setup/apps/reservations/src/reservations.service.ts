import { Injectable } from '@nestjs/common';
import { ReservationRepository } from './reservations.repository';
import { CreateReservationPayload } from './reservations.validations';

@Injectable()
export class ReservationsService {
  constructor(private readonly reservationRepository: ReservationRepository) {}

  createReservation(createReservationPayload: CreateReservationPayload) {
    return this.reservationRepository.create({
      ...createReservationPayload,
      timestamp: new Date(),
      userId: '123',
    });
  }
}
