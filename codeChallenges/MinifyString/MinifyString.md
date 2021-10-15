# Minify a String by Language Type

## Problem:

- When we code, we use lots of unneeded spaces and newlines to make it human readable
- If we want to store or compile, to save space we need to minify the code by removing spaces and newlines
- Careful on what spaces to strip, because some languages have specific syntax rules that break without spaces
- Generally you can strip comments, but sometimes special formatted comments exist that shouldn't be stripped

## Solution:

- Write a function to minify a given string from a given language

## You can use any language, but we prefer:

- Javascript/Typescript
- SQL (SQL-Server a plus)
- PHP
- C#
- Python

## Example:

- Given this JS file (as a string):

```js
const eBacon = {
  kosher: true,
  flavor: "10/10",
};
```

- ...output it in this format:

<!-- prettier-ignore-start -->
```js
const eBacon={kosher:true,flavor:"10/10"}
```
<!-- prettier-ignore-end -->

## Bonus Points:

> Put an `x` inside the brackets to mark them!

- [ ] Create a method that accepts a language type arg:

```js
const minifytext = (type, string) => {
  if (type === "js") {
    // code...
  }
  // etc...
};
```

- Minify these snippets:
- [ ] JS - minify + remove comments

```js
// remove me!

/**
 * don't remove me though!
 * @type {Object}
 */
const eBacon = {
  kosher: true,
  flavor: "10/10",
};

/*
  get me outa here!
*/
```

- [ ] PHP - minify
- [ ] PHP - minify + remove comments

```php
// remove me!

/**
 * don't remove me though!
 * @author Jack
 */
$eBacon = [
  "kosher" => true,
  "flavor" => "10/10",
];

/*
  get me outa here!
*/
```

- [ ] SQL - minify
- [ ] SQL - minify + remove comments

```sql
-- don't mutate the string!
DECLARE eBacon VARCHAR(MAX) = '{
  kosher: true,
  flavor: "10/10",
};'

DECLARE @Bank TABLE (
  id   INT IDENTITY,
  Cash MONEY
)

SELECT Cash
FROM #Bank
WHERE   id IS NOT NULL
    AND Cash >= 100 / 1

/*
  retain valid SQL!
*/
```

- [ ] HTML - minify while retaining valid syntax
- [ ] HTML - minify + remove comments, but not the `if` statements in the head

<!-- prettier-ignore-start -->
```html
<!DOCTYPE html>
<html>
<head>
  <style type="text/css">
    /* ================== */
    /* Generic CSS Resets */
    /* ================== */
    <!--[if mso]>
      div, td {
        padding: 0;
      }

      div {
        margin: 0;
      }
    <![endif]-->

    /* =============== */
    /* Generic Classes */
    /* =============== */
    .some-class {
      width: 50%;
    }
    .someOtherClass {
      width: 80%;
    }
  </style>
</head>
<body>
  <div>
    <table>
      <tbody>
        <tr>
          <!-- Header -->
          <td class="some-class"><h1>Header</h1></td>
          <!-- Content -->
          <td class="someOtherClass"><p>Content</p></td>
        </tr>
      </tbody>
    </table>
  </div>
</body>
</html>
```
<!-- prettier-ignore-end -->

## Templates:

- Javascript:

```js
const minify = (raw) => {
  // Code here...
};

const raw = `
  const eBacon = {
    kosher: true,
    flavor: "10/10",
  };
`;

const minifiedStr = minify(raw);
console.log(minifiedStr);
```

- SQL:

```sql
CREATE FUNCTION [dbo].[MinifyString] (
    @String VARCHAR(MAX)
)
RETURNS VARCHAR(MAX)
BEGIN
    -- Code here...
END

GO

DECLARE @RawString VARCHAR(MAX) = '
  const eBacon = {
    kosher: true,
    flavor: "10/10",
  };
'

SELECT dbo.MinifyString(@RawString)
```

- PHP:

```php
function minify(string $raw)
{
    // Code here...
}

$raw = '
  const eBacon = {
    kosher: true,
    flavor: "10/10",
  };
';

$minifiedStr = formatHTML($raw);
echo $minifiedStr;
```

- C#:

```cs
public class MinifyString
{
    private string minifiedStr;

    public MinifyString(string rawString)
    {
        // Code here...
    }
}

class GetMinifiedStr
{
    private string rawString = @"
      const eBacon = {
        kosher: true,
        flavor: ""10/10"",
      };
    ";

    static void Main()
    {
        var minified = new MinifyString(rawString);
        Console.WriteLine(minified.minifiedStr);
    }
}
```

- Python:

```py
class FormatString:
    def minify(self, raw: str) -> str:
        # Code here...

raw = """
  const eBacon = {
    kosher: true,
    flavor: "10/10",
  };
"""

formatter = FormatString()
minifiedStr = formatter.minify(raw)
print(minifiedStr)
```
