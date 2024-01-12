// automatically generated by the FlatBuffers compiler, do not modify

import * as flatbuffers from 'flatbuffers';

export class SplitPayload {
  bb: flatbuffers.ByteBuffer|null = null;
  bb_pos = 0;
  __init(i:number, bb:flatbuffers.ByteBuffer):SplitPayload {
  this.bb_pos = i;
  this.bb = bb;
  return this;
}

static getRootAsSplitPayload(bb:flatbuffers.ByteBuffer, obj?:SplitPayload):SplitPayload {
  return (obj || new SplitPayload()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

static getSizePrefixedRootAsSplitPayload(bb:flatbuffers.ByteBuffer, obj?:SplitPayload):SplitPayload {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new SplitPayload()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

threshold():number {
  const offset = this.bb!.__offset(this.bb_pos, 4);
  return offset ? this.bb!.readInt32(this.bb_pos + offset) : 0;
}

query():string|null
query(optionalEncoding:flatbuffers.Encoding):string|Uint8Array|null
query(optionalEncoding?:any):string|Uint8Array|null {
  const offset = this.bb!.__offset(this.bb_pos, 6);
  return offset ? this.bb!.__string(this.bb_pos + offset, optionalEncoding) : null;
}

static startSplitPayload(builder:flatbuffers.Builder) {
  builder.startObject(2);
}

static addThreshold(builder:flatbuffers.Builder, threshold:number) {
  builder.addFieldInt32(0, threshold, 0);
}

static addQuery(builder:flatbuffers.Builder, queryOffset:flatbuffers.Offset) {
  builder.addFieldOffset(1, queryOffset, 0);
}

static endSplitPayload(builder:flatbuffers.Builder):flatbuffers.Offset {
  const offset = builder.endObject();
  return offset;
}

static createSplitPayload(builder:flatbuffers.Builder, threshold:number, queryOffset:flatbuffers.Offset):flatbuffers.Offset {
  SplitPayload.startSplitPayload(builder);
  SplitPayload.addThreshold(builder, threshold);
  SplitPayload.addQuery(builder, queryOffset);
  return SplitPayload.endSplitPayload(builder);
}
}
