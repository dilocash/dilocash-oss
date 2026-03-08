/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import { expect, test } from 'vitest'
import fs from 'fs'
import es from '../../i18n/locales/es.json'

test('check auth i18n usage', () => {
    checkI18NForComponents('./components/auth')
})

test('check main i18n usage', () => {
    checkI18NForComponents('./components/main')
})

function checkI18NForComponents(folder) {
    const dirents = fs.readdirSync(folder, { withFileTypes: true });
    dirents.forEach(dirent => {
        if (dirent.isDirectory()) {
            checkI18NForComponents(dirent.parentPath + '/' + dirent.name)
        } else {
            const fileContent = fs.readFileSync(dirent.parentPath + '/' + dirent.name, 'utf-8')
            expect(fileContent).toBeDefined()
            const usedEntries = fileContent.match(/{t\(["']([^"']+)["']\)}/g)
            if (usedEntries) {
                usedEntries.map((entry) => {
                    return entry.replace(/{t\(["']([^"']+)["']\)}/g, '$1')
                }).forEach((entry) => {
                    expect(getValueFromPath(entry, es), entry).toBeDefined()
                })
            }
        }
    })
}

function getValueFromPath(path, obj) {
    const keys = path.split('.')
    let value = obj
    keys.forEach((key) => {
        value = value[key]
    })
    return value
}