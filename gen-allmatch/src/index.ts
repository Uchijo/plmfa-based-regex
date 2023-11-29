import * as fs from 'fs';

type Test = {
    RawPattern: string;
    StrippedPattern: string;
    PositiveExamples: string[];
    NegativeExamples: string[];
    CanHandle: boolean;
    CanProcess: boolean;
}

(async () => {
    const fileContent = fs.readFileSync("tests.json", 'utf8');
    const allTests: Test[] = JSON.parse(fileContent);
    const onlyProcessible = allTests.filter((elem) => elem.CanProcess && !elem.CanHandle)
    const converted = onlyProcessible.map((elem) => {
        try {
            const regexp = RegExp("^" + elem.StrippedPattern + "$")
            const positives = elem.PositiveExamples.map((phrase) => {
                const matches = phrase.match(regexp)
                const pos = matches?.filter((el) => regexp.test(el))
                return pos ? pos : []
            }).flat()
            elem.PositiveExamples = positives
            elem.NegativeExamples = []
            return elem
        } catch (e) {
            return elem
        }
    })
    console.log(JSON.stringify(converted))
})()