import { AbstractDocument } from '@app/common/database/abstract.schema';
import { Prop, SchemaFactory } from '@nestjs/mongoose';

export class ReservationDocument extends AbstractDocument {
  @Prop()
  startDate: Date;

  @Prop()
  endDate: Date;

  @Prop()
  userId: string;

  @Prop()
  placeId: string;

  @Prop()
  invoiceId: string;

  @Prop()
  timestamp: Date;
}

export const ReservationSchema =
  SchemaFactory.createForClass(ReservationDocument);
