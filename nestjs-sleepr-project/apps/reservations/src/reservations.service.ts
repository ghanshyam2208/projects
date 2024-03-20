import { Injectable } from '@nestjs/common';
import { CreateReservationDto } from './dto/CreateReservationDto';
import { ReservationsRepository } from './reservations.repository';
import { UpdateReservationDto } from './dto/UpdateReservationDto';

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

  findAll() {
    return this.reservationsRepository.find({});
  }

  findOne(_id: string) {
    return this.reservationsRepository.findOne({ _id });
  }

  update(_id: string, updateReservationDto: UpdateReservationDto) {
    return this.reservationsRepository.findOneAndUpdate(
      { _id },
      {
        $set: updateReservationDto,
      },
    );
  }

  remove(_id: string) {
    return this.reservationsRepository.findOneAndDelete({ _id });
  }
}
