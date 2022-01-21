# Sort a List of Data Into a Digestible Shape

## Problem:

- The database sends back large objects to the server, and we need to format it to a desired shape for the client
- The frontend needs select pieces of the data to load into a `<Select>` component
- The component in question takes in `options` of this shape:

```js
const options = [
  {
    label: "",
    value: "",
  },
];
```

## Solution:

- Write a function to format the data into a desired shape for the frontend

## You can use one of the following languages:

- Javascript/Typescript
- PHP

## Example:

- Given this data object:

```js
const data = [
  {
    type: Folder,
    id: 1,
    name: 426,
    decription: "Benefits",
    info: "",
    info2: "",
  },
  {
    type: People,
    id: 121,
    name: "John Doe",
    decription: "",
    info: "",
    info2: "",
  },
  {
    type: People,
    id: 128,
    name: "Alex Tato",
    decription: "",
    info: "",
    info2: "",
  },
];
```

```php
$data = [
  [
    "type" => "Folder",
    "id" => 1,
    "name" => 426,
    "decription" => "Benefits",
    "info" => "",
    "info2" => "",
  ],
  [
    "type" => "People",
    "id" => 121,
    "name" => "John Doe",
    "decription" => "",
    "info" => "",
    "info2" => "",
  ],
  [
    "type" => "People",
    "id" => 128,
    "name" => "Alex Tato",
    "decription" => "",
    "info" => "",
    "info2" => "",
  ],
];
```

- ...output it in this format:

```js
const data = {
  Folder: [
    {
      label: "Benefits",
      value: 426,
    },
  ],
  People: [
    {
      label: "John Doe",
      value: 121,
    },
    {
      label: "Alex Tato",
      value: 128,
    },
  ]
};
```

```php
$data = [
  "Folder" => [
    [
      "label" => "Benefits",
      "value" => 426,
    ],
  ],
  "People" => [
    [
      "label" => "John Doe",
      "value" => 121,
    ],
    [
      "label" => "Alex Tato",
      "value" => 128,
    ],
  ]
];
```

## Bonus Points — `Folder` type-specific options and groups

- For `Folder`, the `name` field sometimes hase slashes: `Benefits/Backend`
  - **Note**: this doesn't apply to other option types!
- This denotes "parent-" and "sub-" folders, which we can represent in the data output as groups:

```js
const data = [
  {
    type: Folder,
    id: 1,
    name: 517,
    name: "Benefits",
    info: "",
    info2: "",
  },
  {
    type: Folder,
    id: 1,
    name: 519,
    name: "Benefits/Backend",
    info: "",
    info2: "",
  },
];

// ...will output to...

const options = {
  Folder: [
    {
      label: "Benefits",
      options: [ // this property implies a "group"
        {
          label: "Benefits",
          value: 517,
        },
        {
          label: "Backend",
          value: 519,
        },
      ]
    }
  ],
};
```

```php
$data = [
  [
    "type" => "Folder",
    "id" => 1,
    "name" => 517,
    "name" => "Benefits",
    "info" => "",
    "info2" => "",
  ],
  [
    "type" => "Folder",
    "id" => 1,
    "name" => 519,
    "name" => "Benefits/Backend",
    "info" => "",
    "info2" => "",
  ],
];

// ...will output to...

$options = [
  "Folder" => [
    "label" => "Benefits",
    "options" => [ // this property implies a "group"
      [
        "label" => "Benefits",
        "value" => 517,
      ],
      [
        "label" => "Backend",
        "value" => 519,
      ],
    ],
  ],
];
```

## Templates:

- Javascript:

```js
const sortData = (rawData) => {
  // Code here...
};

// import mock data or paste it here

const optionsArray = sortData(rawData);
console.log(optionsArray);
```

- PHP:

```php
function sortData(array $rawData)
{
    // Code here...
}

// paste mock data here
// $rawData =

$optionsArray = sortData($rawData);
echo $optionsArray;
```
