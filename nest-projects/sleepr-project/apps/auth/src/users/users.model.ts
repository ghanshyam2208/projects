import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { AbstractDocument } from '@app/common/database/abstract.schema';

@Schema({ versionKey: false })
export class UsersDocument extends AbstractDocument {
  @Prop()
  email: string;

  @Prop()
  password: string;
}

export const UsersSchema = SchemaFactory.createForClass(UsersDocument);
