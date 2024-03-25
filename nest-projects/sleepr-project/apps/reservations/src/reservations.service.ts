import { Injectable } from '@nestjs/common';
import { ReservationsRepository } from './reservations.repository';
import { ReservationsDocument } from './model/reservations.model';
import { CreateReservationDto } from './dto/CreateReservationDto';
import { UpdateReservationDto } from './dto/UpdateReservationDto';

@Injectable()
export class ReservationsService {
  constructor(
    private readonly reservationsRepository: ReservationsRepository,
  ) {}
  findAll(): Promise<ReservationsDocument[]> {
    return this.reservationsRepository.find({});
  }
  createReservation(createReservationDto: CreateReservationDto) {
    return this.reservationsRepository.create({
      ...createReservationDto,
      userId: '123',
      timestamp: new Date(),
    });
  }

  findOne(_id: string) {
    return this.reservationsRepository.findOne({
      _id,
    });
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
