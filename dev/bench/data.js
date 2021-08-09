window.BENCHMARK_DATA = {
  "lastUpdate": 1628539842067,
  "repoUrl": "https://github.com/romdo/go-conventionalcommit",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "contact@jimeh.me",
            "name": "Jim Myhrberg",
            "username": "jimeh"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "e8ca0094405c69ca11e814cf4c3f6fa4e29fd964",
          "message": "Merge pull request #1 from romdo/add-rawmessage\n\nfeat(parser): implement RawMessage",
          "timestamp": "2021-08-09T21:06:35+01:00",
          "tree_id": "2be1c5766f33a3d9786adcc8f0e407f71ca17df2",
          "url": "https://github.com/romdo/go-conventionalcommit/commit/e8ca0094405c69ca11e814cf4c3f6fa4e29fd964"
        },
        "date": 1628539841336,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkLines_Bytes__single_line",
            "value": 34.5,
            "unit": "ns/op",
            "extra": "33146340 times"
          },
          {
            "name": "BenchmarkLines_Bytes__single_line_with_trailing_LF",
            "value": 40.6,
            "unit": "ns/op",
            "extra": "29882890 times"
          },
          {
            "name": "BenchmarkLines_Bytes__single_line_with_trailing_CRLF",
            "value": 42,
            "unit": "ns/op",
            "extra": "29006472 times"
          },
          {
            "name": "BenchmarkLines_Bytes__single_line_with_trailing_CR",
            "value": 40.8,
            "unit": "ns/op",
            "extra": "29639756 times"
          },
          {
            "name": "BenchmarkLines_Bytes__multi_line_separated_by_LF",
            "value": 137,
            "unit": "ns/op",
            "extra": "8818696 times"
          },
          {
            "name": "BenchmarkLines_Bytes__multi_line_separated_by_CRLF",
            "value": 142,
            "unit": "ns/op",
            "extra": "8162557 times"
          },
          {
            "name": "BenchmarkLines_Bytes__multi_line_separated_by_CR",
            "value": 136,
            "unit": "ns/op",
            "extra": "8745592 times"
          },
          {
            "name": "BenchmarkLines_String__single_line",
            "value": 37.9,
            "unit": "ns/op",
            "extra": "30263858 times"
          },
          {
            "name": "BenchmarkLines_String__single_line_with_trailing_LF",
            "value": 46.3,
            "unit": "ns/op",
            "extra": "26519558 times"
          },
          {
            "name": "BenchmarkLines_String__single_line_with_trailing_CRLF",
            "value": 46.9,
            "unit": "ns/op",
            "extra": "26067218 times"
          },
          {
            "name": "BenchmarkLines_String__single_line_with_trailing_CR",
            "value": 45.8,
            "unit": "ns/op",
            "extra": "26700357 times"
          },
          {
            "name": "BenchmarkLines_String__multi_line_separated_by_LF",
            "value": 194,
            "unit": "ns/op",
            "extra": "6139024 times"
          },
          {
            "name": "BenchmarkLines_String__multi_line_separated_by_CRLF",
            "value": 204,
            "unit": "ns/op",
            "extra": "5944452 times"
          },
          {
            "name": "BenchmarkLines_String__multi_line_separated_by_CR",
            "value": 198,
            "unit": "ns/op",
            "extra": "6105099 times"
          },
          {
            "name": "BenchmarkNewRawMessage__nil",
            "value": 42.4,
            "unit": "ns/op",
            "extra": "27482770 times"
          },
          {
            "name": "BenchmarkNewRawMessage__empty",
            "value": 43,
            "unit": "ns/op",
            "extra": "28105476 times"
          },
          {
            "name": "BenchmarkNewRawMessage__single_space",
            "value": 176,
            "unit": "ns/op",
            "extra": "6438363 times"
          },
          {
            "name": "BenchmarkNewRawMessage__subject_only",
            "value": 266,
            "unit": "ns/op",
            "extra": "4322313 times"
          },
          {
            "name": "BenchmarkNewRawMessage__subject_and_body",
            "value": 859,
            "unit": "ns/op",
            "extra": "1327281 times"
          },
          {
            "name": "BenchmarkNewRawMessage__subject_and_body_with_CRLF_line_breaks",
            "value": 864,
            "unit": "ns/op",
            "extra": "1362050 times"
          },
          {
            "name": "BenchmarkNewRawMessage__subject_and_body_with_CR_line_breaks",
            "value": 876,
            "unit": "ns/op",
            "extra": "1381249 times"
          },
          {
            "name": "BenchmarkNewRawMessage__separated_by_whitespace_line",
            "value": 881,
            "unit": "ns/op",
            "extra": "1381087 times"
          },
          {
            "name": "BenchmarkNewRawMessage__subject_and_long_body",
            "value": 5833,
            "unit": "ns/op",
            "extra": "208868 times"
          },
          {
            "name": "BenchmarkRawMessage_Bytes__empty",
            "value": 8.47,
            "unit": "ns/op",
            "extra": "142406499 times"
          },
          {
            "name": "BenchmarkRawMessage_Bytes__single_space",
            "value": 24.6,
            "unit": "ns/op",
            "extra": "48954320 times"
          },
          {
            "name": "BenchmarkRawMessage_Bytes__subject_only",
            "value": 37.4,
            "unit": "ns/op",
            "extra": "29611442 times"
          },
          {
            "name": "BenchmarkRawMessage_Bytes__subject_and_body",
            "value": 56.4,
            "unit": "ns/op",
            "extra": "21679516 times"
          },
          {
            "name": "BenchmarkRawMessage_Bytes__subject_and_body_with_CRLF_line_breaks",
            "value": 60.6,
            "unit": "ns/op",
            "extra": "19804483 times"
          },
          {
            "name": "BenchmarkRawMessage_Bytes__subject_and_body_with_CR_line_breaks",
            "value": 55.3,
            "unit": "ns/op",
            "extra": "21668710 times"
          },
          {
            "name": "BenchmarkRawMessage_Bytes__separated_by_whitespace_line",
            "value": 57,
            "unit": "ns/op",
            "extra": "21257516 times"
          },
          {
            "name": "BenchmarkRawMessage_Bytes__subject_and_long_body",
            "value": 447,
            "unit": "ns/op",
            "extra": "2630110 times"
          },
          {
            "name": "BenchmarkRawMessage_String__empty",
            "value": 10.8,
            "unit": "ns/op",
            "extra": "100000000 times"
          },
          {
            "name": "BenchmarkRawMessage_String__single_space",
            "value": 27.2,
            "unit": "ns/op",
            "extra": "43455600 times"
          },
          {
            "name": "BenchmarkRawMessage_String__subject_only",
            "value": 42.4,
            "unit": "ns/op",
            "extra": "27122796 times"
          },
          {
            "name": "BenchmarkRawMessage_String__subject_and_body",
            "value": 88.1,
            "unit": "ns/op",
            "extra": "13643185 times"
          },
          {
            "name": "BenchmarkRawMessage_String__subject_and_body_with_CRLF_line_breaks",
            "value": 88.9,
            "unit": "ns/op",
            "extra": "13576803 times"
          },
          {
            "name": "BenchmarkRawMessage_String__subject_and_body_with_CR_line_breaks",
            "value": 87.8,
            "unit": "ns/op",
            "extra": "13678256 times"
          },
          {
            "name": "BenchmarkRawMessage_String__separated_by_whitespace_line",
            "value": 90.3,
            "unit": "ns/op",
            "extra": "12800388 times"
          },
          {
            "name": "BenchmarkRawMessage_String__subject_and_long_body",
            "value": 693,
            "unit": "ns/op",
            "extra": "1724236 times"
          }
        ]
      }
    ]
  }
}