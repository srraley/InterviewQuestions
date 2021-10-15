# Format an HTML string into valid email HTML

## Problem:

- You can embed raw HTML into emails, but some email clients have very poor feature support
- The feature in question you must fix is moving styles in the `<head>` into in-line styles

## Solution:

- Write a function that takes an HTML string and returns a new string with any styles in the `<head>` moved to inline styles

## You can use any language, but we prefer:

- Javascript/Typescript
- SQL (SQL-Server a plus)
- PHP
- C#
- Python

## Example:

- Given this HTML (as a string):

```html
<head>
  <style>
    .some-class {
      width: 50%;
    }
    .someOtherClass {
      width: 80%;
      height: 20px;
    }
  </style>
</head>
<body>
  <div>
    <table>
      <tbody>
        <tr>
          <td class="some-class"><h1>Header</h1></td>
          <td class="someOtherClass"><p>Content</p></td>
        </tr>
      </tbody>
    </table>
  </div>
</body>
```

- ... output it in this format:

```html
<head>
  <style>
    .some-class {
      width: 50%;
    }
    .someOtherClass {
      width: 80%;
      height: 20px;
    }
  </style>
</head>
<body>
  <div>
    <table>
      <tbody>
        <tr>
          <td class="some-class" style="width: 50%;"><h1>Header</h1></td>
          <td class="someOtherClass" style="width: 80%; height: 20px;"><p>Content</p></td>
        </tr>
      </tbody>
    </table>
  </div>
</body>
```

## Bonus Points:

> Put an `x` inside the brackets to mark them!

- Reduce the length of the string by:
  - [ ] Remove styles from the head
  - [ ] Remove class names on elements
  - [ ] Remove excess spaces and newlines while retaining valid HTML
- [ ] Add this to the stylesheet in the head:

```html
<style>
  table {
    width: 100%;
    margin: 0;
    text-align: center;
    border-collapse: collapse;
    border-spacing: 0;
  }
  .table-container {
    width: 100%;
    margin: 0 auto;
    text-align: center;
    border-collapse: collapse;
    border-spacing: 0;
    background-color: #ffffff;
  }
</style>
```

- [ ] Add this below the "Content" `<td>`:

```html
<tr>
  <td class="content-td"><p class="content-td-p">Content2</p></td>
</tr>
```

## Templates:

- Javascript:

```js
const formatHTML = (rawHTML) => {
  // Code here...
};

const rawHTML = `
  <!DOCTYPE html>
  <html>
  <head>
    <style>
      .header-td {
        width: 100%;
        border-radius: 4px 4px 0 0;
      }
      .header-td-h1 {
        font-weight: 500;
        text-align: center;
        font-size: 20px;
      }
      .content-td {
        width: 100%;
        padding: 20px 15px;
      }
      .content-td-p {
        text-align: center;
      }
      .footer {
        margin: 0 auto;
        width: fit-content;
        border-collapse: unset;
      }
      .header-td {
        shit: YES!;
      }
      .footer-td {
        border-top: 1px solid #0066cc;
        padding: 12px 30px;
        text-align: center;
        color: #4c4c4c;
        font-size: 12px;
        line-height: 18px;
      }
    </style>
  </head>
  <body>
    <div>
      <table class="table-container">
        <tbody>
          <tr>
            <td class="header-td"><h1 class="header-td-h1">Header</h1></td>
          </tr>
          <tr>
            <td class="content-td"><p class="content-td-p">Content</p></td>
          </tr>
        </tbody>
      </table>
      <table class="footer">
        <tbody>
          <tr>
            <td class="footer-td">
              Regards,
              <br />
              <strong>eBacon Team</strong>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </body>
  </html>
`;

const emailHTML = formatHTML(rawHTML);
console.log(emailHTML);
```

- SQL:

```sql
DROP FUNCTION IF EXISTS [dbo].[FormatHTML]
GO

CREATE FUNCTION [dbo].[FormatHTML] (
    @HTML VARCHAR(MAX)
)
RETURNS VARCHAR(MAX)
BEGIN
    -- Code here...
END

GO

DECLARE @RawHTML VARCHAR(MAX) = '
  <!DOCTYPE html>
  <html>
  <head>
    <style>
      .header-td {
        width: 100%;
        border-radius: 4px 4px 0 0;
      }
      .header-td-h1 {
        font-weight: 500;
        text-align: center;
        font-size: 20px;
      }
      .content-td {
        width: 100%;
        padding: 20px 15px;
      }
      .content-td-p {
        text-align: center;
      }
      .footer {
        margin: 0 auto;
        width: fit-content;
        border-collapse: unset;
      }
      .header-td {
        shit: YES!;
      }
      .footer-td {
        border-top: 1px solid #0066cc;
        padding: 12px 30px;
        text-align: center;
        color: #4c4c4c;
        font-size: 12px;
        line-height: 18px;
      }
    </style>
  </head>
  <body>
    <div>
      <table class="table-container">
        <tbody>
          <tr>
            <td class="header-td"><h1 class="header-td-h1">Header</h1></td>
          </tr>
          <tr>
            <td class="content-td"><p class="content-td-p">Content</p></td>
          </tr>
        </tbody>
      </table>
      <table class="footer">
        <tbody>
          <tr>
            <td class="footer-td">
              Regards,
              <br />
              <strong>eBacon Team</strong>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </body>
  </html>
'

SELECT dbo.FormatHTML(@RawHTML)
```

- PHP:

```php
function formatHTML(string $raw)
{
    // Code here...
}

$rawHTML = '
  <!DOCTYPE html>
  <html>
  <head>
    <style>
      .header-td {
        width: 100%;
        border-radius: 4px 4px 0 0;
      }
      .header-td-h1 {
        font-weight: 500;
        text-align: center;
        font-size: 20px;
      }
      .content-td {
        width: 100%;
        padding: 20px 15px;
      }
      .content-td-p {
        text-align: center;
      }
      .footer {
        margin: 0 auto;
        width: fit-content;
        border-collapse: unset;
      }
      .header-td {
        shit: YES!;
      }
      .footer-td {
        border-top: 1px solid #0066cc;
        padding: 12px 30px;
        text-align: center;
        color: #4c4c4c;
        font-size: 12px;
        line-height: 18px;
      }
    </style>
  </head>
  <body>
    <div>
      <table class="table-container">
        <tbody>
          <tr>
            <td class="header-td"><h1 class="header-td-h1">Header</h1></td>
          </tr>
          <tr>
            <td class="content-td"><p class="content-td-p">Content</p></td>
          </tr>
        </tbody>
      </table>
      <table class="footer">
        <tbody>
          <tr>
            <td class="footer-td">
              Regards,
              <br />
              <strong>eBacon Team</strong>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </body>
  </html>
';

$emailHTML = formatHTML($rawHTML);
echo $emailHTML;
```

- C#:

```cs
public class FormatEmail
{
    private string formattedHTML;

    public FormatEmail(string rawHTML)
    {
        // Code here...
    }
}

class GetEmailHTML
{
    private string rawHTML = @"
      <!DOCTYPE html>
      <html>
      <head>
        <style>
          .header-td {
            width: 100%;
            border-radius: 4px 4px 0 0;
          }
          .header-td-h1 {
            font-weight: 500;
            text-align: center;
            font-size: 20px;
          }
          .content-td {
            width: 100%;
            padding: 20px 15px;
          }
          .content-td-p {
            text-align: center;
          }
          .footer {
            margin: 0 auto;
            width: fit-content;
            border-collapse: unset;
          }
          .header-td {
            shit: YES!;
          }
          .footer-td {
            border-top: 1px solid #0066cc;
            padding: 12px 30px;
            text-align: center;
            color: #4c4c4c;
            font-size: 12px;
            line-height: 18px;
          }
        </style>
      </head>
      <body>
        <div>
          <table class=""table-container"">
            <tbody>
              <tr>
                <td class=""header-td""><h1 class=""header-td-h1"">Header</h1></td>
              </tr>
              <tr>
                <td class=""content-td""><p class=""content-td-p"">Content</p></td>
              </tr>
            </tbody>
          </table>
          <table class=""footer"">
            <tbody>
              <tr>
                <td class=""footer-td"">
                  Regards,
                  <br />
                  <strong>eBacon Team</strong>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </body>
      </html>
    ";

    static void Main()
    {
        var email = new FormatEmail(rawHTML);
        Console.WriteLine(email.formattedHTML);
    }
}
```

- Python:

```py
class HTMLFormatter:
    def email(self, rawHTML: str) -> str:
        # Code here...

rawHTML = """
  <!DOCTYPE html>
  <html>
  <head>
    <style>
      .header-td {
        width: 100%;
        border-radius: 4px 4px 0 0;
      }
      .header-td-h1 {
        font-weight: 500;
        text-align: center;
        font-size: 20px;
      }
      .content-td {
        width: 100%;
        padding: 20px 15px;
      }
      .content-td-p {
        text-align: center;
      }
      .footer {
        margin: 0 auto;
        width: fit-content;
        border-collapse: unset;
      }
      .header-td {
        shit: YES!;
      }
      .footer-td {
        border-top: 1px solid #0066cc;
        padding: 12px 30px;
        text-align: center;
        color: #4c4c4c;
        font-size: 12px;
        line-height: 18px;
      }
    </style>
  </head>
  <body>
    <div>
      <table class="table-container">
        <tbody>
          <tr>
            <td class="header-td"><h1 class="header-td-h1">Header</h1></td>
          </tr>
          <tr>
            <td class="content-td"><p class="content-td-p">Content</p></td>
          </tr>
        </tbody>
      </table>
      <table class="footer">
        <tbody>
          <tr>
            <td class="footer-td">
              Regards,
              <br />
              <strong>eBacon Team</strong>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </body>
  </html>
"""

formatter = HTMLFormatter()
emailHTML = formatter.email(rawHTML)
print(emailHTML)
```
