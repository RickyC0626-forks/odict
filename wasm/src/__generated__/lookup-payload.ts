// automatically generated by the FlatBuffers compiler, do not modify

import * as flatbuffers from 'flatbuffers';

import { MarkdownStrategy } from './markdown-strategy.js';


export class LookupPayload {
  bb: flatbuffers.ByteBuffer|null = null;
  bb_pos = 0;
  __init(i:number, bb:flatbuffers.ByteBuffer):LookupPayload {
  this.bb_pos = i;
  this.bb = bb;
  return this;
}

static getRootAsLookupPayload(bb:flatbuffers.ByteBuffer, obj?:LookupPayload):LookupPayload {
  return (obj || new LookupPayload()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

static getSizePrefixedRootAsLookupPayload(bb:flatbuffers.ByteBuffer, obj?:LookupPayload):LookupPayload {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new LookupPayload()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

follow():boolean {
  const offset = this.bb!.__offset(this.bb_pos, 4);
  return offset ? !!this.bb!.readInt8(this.bb_pos + offset) : false;
}

split():number {
  const offset = this.bb!.__offset(this.bb_pos, 6);
  return offset ? this.bb!.readInt32(this.bb_pos + offset) : 0;
}

markdown():MarkdownStrategy {
  const offset = this.bb!.__offset(this.bb_pos, 8);
  return offset ? this.bb!.readInt16(this.bb_pos + offset) : MarkdownStrategy.Disable;
}

queries(index: number):string
queries(index: number,optionalEncoding:flatbuffers.Encoding):string|Uint8Array
queries(index: number,optionalEncoding?:any):string|Uint8Array|null {
  const offset = this.bb!.__offset(this.bb_pos, 10);
  return offset ? this.bb!.__string(this.bb!.__vector(this.bb_pos + offset) + index * 4, optionalEncoding) : null;
}

queriesLength():number {
  const offset = this.bb!.__offset(this.bb_pos, 10);
  return offset ? this.bb!.__vector_len(this.bb_pos + offset) : 0;
}

static startLookupPayload(builder:flatbuffers.Builder) {
  builder.startObject(4);
}

static addFollow(builder:flatbuffers.Builder, follow:boolean) {
  builder.addFieldInt8(0, +follow, +false);
}

static addSplit(builder:flatbuffers.Builder, split:number) {
  builder.addFieldInt32(1, split, 0);
}

static addMarkdown(builder:flatbuffers.Builder, markdown:MarkdownStrategy) {
  builder.addFieldInt16(2, markdown, MarkdownStrategy.Disable);
}

static addQueries(builder:flatbuffers.Builder, queriesOffset:flatbuffers.Offset) {
  builder.addFieldOffset(3, queriesOffset, 0);
}

static createQueriesVector(builder:flatbuffers.Builder, data:flatbuffers.Offset[]):flatbuffers.Offset {
  builder.startVector(4, data.length, 4);
  for (let i = data.length - 1; i >= 0; i--) {
    builder.addOffset(data[i]!);
  }
  return builder.endVector();
}

static startQueriesVector(builder:flatbuffers.Builder, numElems:number) {
  builder.startVector(4, numElems, 4);
}

static endLookupPayload(builder:flatbuffers.Builder):flatbuffers.Offset {
  const offset = builder.endObject();
  return offset;
}

static createLookupPayload(builder:flatbuffers.Builder, follow:boolean, split:number, markdown:MarkdownStrategy, queriesOffset:flatbuffers.Offset):flatbuffers.Offset {
  LookupPayload.startLookupPayload(builder);
  LookupPayload.addFollow(builder, follow);
  LookupPayload.addSplit(builder, split);
  LookupPayload.addMarkdown(builder, markdown);
  LookupPayload.addQueries(builder, queriesOffset);
  return LookupPayload.endLookupPayload(builder);
}
}
