/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import { expect, test } from 'vitest'
import en from '../locales/en.json'
import es from '../locales/es.json'

test('same entries exist in both languages', () => {
    checkAllKeysExist(en, es)
})

function checkAllKeysExist(a, b) {
    expect(Object.keys(a)).toEqual(Object.keys(b))
    for (const key of Object.keys(a)) {
        if (typeof a[key] === 'object' && typeof b[key] === 'object') {
            checkAllKeysExist(a[key], b[key])
        }
    }
}

test('all entries have values in both languages', () => {
    expect(Object.keys(en).length).toEqual(Object.keys(es).length)
    expect(Object.values(en).length).toEqual(Object.values(es).length)
})
