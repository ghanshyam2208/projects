import { AbstractRepository } from '@app/common';
import { ReservationDocument } from './reservation.schema';
import { Injectable, Logger } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

@Injectable()
export class ReservationRepository extends AbstractRepository<ReservationDocument> {
  protected readonly logger = new Logger(ReservationRepository.name);

  constructor(
    @InjectModel(ReservationDocument.name)
    reservationDocument: Model<ReservationDocument>,
  ) {
    super(reservationDocument);
  }
}
