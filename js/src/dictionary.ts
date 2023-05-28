import { Builder } from "flatbuffers";

import { CompilePayload } from "./__generated__/compile-payload.js";
import {
  LookupPayload,
  SearchPayload,
  SplitPayload,
  WritePayload,
} from "./__generated__/index.js";
import { ODictMethod } from "./__generated__/odict-method.js";
import { startService } from "./service.js";
import type {
  DictionaryOptions,
  Entry,
  LookupOptions,
  Query,
  SearchOptions,
} from "./types.js";
import { generateOutputPath, queryToString } from "./utils.js";

function createBuilder() {
  return new Builder(1024);
}

class Dictionary {
  private readonly options: DictionaryOptions;

  constructor(
    public readonly path: string,
    options: Partial<DictionaryOptions> = {}
  ) {
    this.options = options;
  }

  /**
   * Compiles an ODXML file to a compiled dictionary
   *
   * @param xmlPath The path to the XML file to compile
   * @param outPath The destination path (autogenerated if omitted)
   * @returns A pointer to the compiled Dictionary
   */
  static async compile(xmlPath: string, outPath?: string): Promise<Dictionary> {
    const out = outPath ?? generateOutputPath(xmlPath);
    const builder = createBuilder();

    const pathS = builder.createString(xmlPath);
    const outS = builder.createString(out);

    CompilePayload.startCompilePayload(builder);
    CompilePayload.addPath(builder, pathS);
    CompilePayload.addOut(builder, outS);

    const payload = CompilePayload.endCompilePayload(builder);

    builder.finish(payload);

    await startService().run(ODictMethod.Compile, builder.asUint8Array());

    return new Dictionary(out);
  }

  /**
   * Writes an ODXML string to a compiled dictionary file
   *
   * @param xml The XML to compile
   * @param outPath The destination path
   * @returns A pointer to the compiled dictionary
   */
  static async write(xml: string, outPath: string): Promise<Dictionary> {
    const builder = createBuilder();
    const xmlS = builder.createString(xml);
    const outS = builder.createString(outPath);

    WritePayload.startWritePayload(builder);
    WritePayload.addXml(builder, xmlS);
    WritePayload.addOut(builder, outS);

    const payload = WritePayload.endWritePayload(builder);

    builder.finish(payload);

    await startService().run(ODictMethod.Write, builder.asUint8Array());

    return new Dictionary(outPath);
  }

  /**
   * Indexes a compiled dictionary so it can be searched via the search() method
   */
  async index() {
    await startService(this.path).run(ODictMethod.Index);
  }

  /**
   * Searches a compiled dictionary using full-text search
   *
   * @param query The query (or a list of queries) to search
   * @param options Search options
   * @returns A nested array of entries
   */
  async search(
    query: Query | Query[],
    options: SearchOptions = {}
  ): Promise<Entry[][]> {
    const queries = Array.isArray(query) ? query : [query];

    return Promise.all(
      queries.map(queryToString).map(async (query) => {
        const builder = createBuilder();
        const queryS = builder.createString(query);

        SearchPayload.startSearchPayload(builder);
        SearchPayload.addQuery(builder, queryS);
        SearchPayload.addExact(builder, options.exact ?? false);
        SearchPayload.addForce(builder, options.force ?? false);

        const payload = SearchPayload.endSearchPayload(builder);

        builder.finish(payload);

        return startService(this.path).run<Entry[]>(
          ODictMethod.Search,
          builder.asUint8Array()
        );
      })
    );
  }

  /**
   * Returns a list of all headwords in the dictionary
   *
   * @returns A list of all headwords in the dictionary
   */
  async lexicon(): Promise<string[]> {
    return startService(this.path).run<string[]>(ODictMethod.Lexicon);
  }

  /**
   * Looks up an entry in the dictionary
   *
   * @param query The query (or a list of queries) to lookup
   * @param options Lookup options
   * @returns A nested array of entries
   */
  async lookup(
    query: Query | Query[],
    options: LookupOptions = {}
  ): Promise<Entry[][]> {
    const queries = Array.isArray(query) ? query : [query];

    const { follow, split = this.options.defaultSplitThreshold } = options;

    const builder = createBuilder();

    const queriesS = queries
      .map(queryToString)
      .map((str) => builder.createString(str));

    const queriesV = LookupPayload.createQueriesVector(builder, queriesS);

    LookupPayload.startLookupPayload(builder);
    LookupPayload.addQueries(builder, queriesV);
    LookupPayload.addFollow(builder, follow ?? false);
    LookupPayload.addSplit(builder, split ?? 0);

    const payload = LookupPayload.endLookupPayload(builder);

    builder.finish(payload);

    return startService(this.path).run<Entry[][]>(
      ODictMethod.Lookup,
      builder.asUint8Array()
    );
  }

  /**
   * Splits a query into its definable subparts
   *
   * @param query The query (or a list of queries) to lookup
   * @param options Lookup options
   * @returns A nested array of entries
   */
  async split(query: string, threshold: number): Promise<Entry[]> {
    const builder = createBuilder();
    const queryS = builder.createString(query);

    SplitPayload.startSplitPayload(builder);
    SplitPayload.addThreshold(builder, threshold);
    SplitPayload.addQuery(builder, queryS);

    const payload = SplitPayload.endSplitPayload(builder);

    builder.finish(payload);

    return startService(this.path).run<Entry[]>(
      ODictMethod.Split,
      builder.asUint8Array()
    );
  }
}

export { Dictionary };
