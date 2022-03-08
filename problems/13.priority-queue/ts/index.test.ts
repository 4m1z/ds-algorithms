import { describe, it, assert } from 'vitest';
import { PriorityQueueImpl } from './PriorityQueue';

describe.skip('priority queue', () => {
    it('making a max heap via', () => {
        const pq = new PriorityQueueImpl([], (a: number, b: number) => a > b);
        const pq2 = new PriorityQueueImpl(
            [25, 40, 15, 20, 10, 50, 35],
            (a: number, b: number) => a > b
        );

        pq.insert(15);
        pq.insert(35);
        pq.insert(10);
        pq.insert(25);
        pq.insert(20);
        pq.insert(50);
        pq.insert(40);

        assert.deepEqual(pq.getHeap(), [50, 25, 40, 15, 20, 10, 35]);
        assert.deepEqual(pq2.getHeap(), [50, 40, 35, 20, 10, 15, 25]);
    });

    it('making a min heap', () => {
        const pq = new PriorityQueueImpl([], (a: number, b: number) => a < b);
        const pq2 = new PriorityQueueImpl(
            [25, 40, 15, 20, 10, 50, 35],
            (a: number, b: number) => a < b
        );

        pq.insert(15);
        pq.insert(35);
        pq.insert(10);
        pq.insert(25);
        pq.insert(20);
        pq.insert(50);
        pq.insert(40);

        assert.deepEqual(pq.getHeap(), [10, 20, 15, 35, 25, 50, 40]);
        assert.deepEqual(pq2.getHeap(), [10, 20, 15, 25, 40, 50, 35]);
    });
});
