/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import React from 'react';
import { cardStyle } from './styles';
import type { VariantProps } from '@gluestack-ui/utils/nativewind-utils';

type ICardProps = React.ComponentPropsWithoutRef<'div'> &
  VariantProps<typeof cardStyle>;

const Card = React.forwardRef<HTMLDivElement, ICardProps>(function Card(
  { className, size = 'md', variant = 'elevated', ...props },
  ref
) {
  return (
    <div
      className={cardStyle({ size, variant, class: className })}
      {...props}
      ref={ref}
    />
  );
});

Card.displayName = 'Card';

export { Card };
