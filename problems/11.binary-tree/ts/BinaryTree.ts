import { Node } from './Node';

export class BinaryTree<T> {
    private root: Node<T> | null = null;

    constructor(root: Node<T>) {
        this.setRoot(root);
    }

    public getRoot(): Node<T> | null {
        return this.root;
    }

    private setRoot(root: Node<T>) {
        this.root = root;
    }

    public findMaxDepth(node: Node<T> | null, max = 0): number {
        if (!node) return max;

        max = max + 1;

        return Math.max.apply(null, [
            this.findMaxDepth(node.left, max),
            this.findMaxDepth(node.right, max),
        ]);
    }

    public findLevelOrderValues(): T[][] {
        let result: T[][] = [];
        let temp = this.root;

        if (!temp) return result;

        let queue = [[temp]];

        while (queue.length > 0) {
            let current = queue.shift() as Node<T>[];
            let level: Node<T>[] = [];
            result.push(
                current.map((node) => {
                    node.left && level.push(node.left);
                    node.right && level.push(node.right);
                    return node.value;
                })
            );

            if (level.length > 0) queue = [level];
        }

        return result;
    }

    public findLevelOrderValuesVisibleFromRight(): T[] {
        let result: T[][] = [];
        let temp = this.root;

        if (!temp) return result.flat();

        let queue = [[temp]];

        while (queue.length > 0) {
            let current = queue.shift() as Node<T>[];
            let level: Node<T>[] = [];
            result.push(
                current.map((node) => {
                    node.left && level.push(node.left);
                    node.right && level.push(node.right);
                    return node.value;
                })
            );

            if (level.length > 0) queue = [level];
        }

        return result.map((item: T[]) => [item[item.length - 1]]).flat();
    }
}
