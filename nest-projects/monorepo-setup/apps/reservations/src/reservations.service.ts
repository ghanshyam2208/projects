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

  getReservations() {
    return this.reservationRepository.find({});
  }

  getReservationById(_id: string) {
    return this.reservationRepository.findOne({ _id });
  }

  updateReservation(
    _id: string,
    createReservationPayload: CreateReservationPayload,
  ) {
    return this.reservationRepository.findOneAndUpdate(
      { _id },
      {
        $set: createReservationPayload,
      },
    );
  }

  deleteReservationById(_id: string) {
    return this.reservationRepository.findOneAndDelete({ _id });
  }
}
