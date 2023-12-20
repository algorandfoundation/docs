[@algorandfoundation/algokit-utils](../index.md) / [types/testing](../modules/types_testing.md) / AlgorandFixture

# Interface: AlgorandFixture

[types/testing](../modules/types_testing.md).AlgorandFixture

An Algorand automated testing fixture

## Table of contents

### Properties

- [beforeEach](types_testing.AlgorandFixture.md#beforeeach)

### Accessors

- [context](types_testing.AlgorandFixture.md#context)

## Properties

### beforeEach

• **beforeEach**: () => `Promise`\<`void`\>

#### Type declaration

▸ (): `Promise`\<`void`\>

Testing framework agnostic handler method to run before each test to prepare the `context` for that test.

##### Returns

`Promise`\<`void`\>

#### Defined in

[src/types/testing.ts:74](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L74)

## Accessors

### context

• `get` **context**(): [`AlgorandTestAutomationContext`](types_testing.AlgorandTestAutomationContext.md)

Retrieve the current context.
Useful with destructuring.

#### Returns

[`AlgorandTestAutomationContext`](types_testing.AlgorandTestAutomationContext.md)

**`Example`**

```typescript
test('My test', () => {
    const {algod, indexer, testAccount, ...} = algorand.context
})
```

#### Defined in

[src/types/testing.ts:69](https://github.com/algorandfoundation/algokit-utils-ts/blob/main/src/types/testing.ts#L69)
