#include "flatbuffers/idl.h"
#include "flatbuffers/util.h"

#include "monster_generated.h" // Already includes "flatbuffers/flatbuffers.h".

using namespace MyGame::Sample;

int main(int argc, const char * argv[]) {
    if (argc != 5) {
        printf("invalid arg number, need [fbs dir] [fb file path] [json in file path] [out file name]");
        return 1;
    }

    std::string schemafile;
    std::string jsonfile;
    bool ok = flatbuffers::LoadFile(argv[2], false, &schemafile) &&
        flatbuffers::LoadFile(argv[3], false, &jsonfile);
    if (!ok) {
        printf("couldn't load fb file path or json file path\n");
        return 1;
    }


    flatbuffers::IDLOptions opts;
    opts.strict_json = true;
    flatbuffers::Parser parser(opts);
    const char *include_directories[] = { argv[1], nullptr };
    ok = parser.Parse(schemafile.c_str(), include_directories);
    if (!ok) {
        printf("couldn't parse the fb schema%s\n", schemafile.c_str());
        return 1;
    }
    ok = parser.Parse(jsonfile.c_str(), include_directories);
    if (!ok) {
        printf("couldn't parse the fb json%s\n", jsonfile.c_str());
        return 1;
    }

    std::string jsongen;
    ok = GenerateText(parser, parser.builder_.GetBufferPointer(), &jsongen);
    if (!ok) {
        printf("Couldn't serialize parsed data to JSON!\n");
        return 1;
    }
    if (jsongen != jsonfile) {
        printf("%s----------------\n%s", jsongen.c_str(), jsonfile.c_str());
    }

    ok = GenerateTextFile(parser, argv[1], argv[4]);
    if (!ok) {
        printf("Couldn't convert fb to json\n");
        return 1;
    }

    ok = GenerateBinary(parser, argv[1], argv[4]);
    if (!ok) {
        printf("Couldn't convert fb to json\n");
        return 1;
    }

    printf("ALL OK!!!!\n");
    return 0;
}
