import * as Joi from 'joi';

export class CreateReservationDto {
  startDate: Date;
  endDate: Date;
  placeId: string;
  invoiceId: string;
}

export const CreateReservationSchema = Joi.object({
  startDate: Joi.date().required(),
  endDate: Joi.date().required(),
  placeId: Joi.string().required(),
  invoiceId: Joi.string().required(),
}).options({
  abortEarly: false,
});
