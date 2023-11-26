# https://leetcode.com/problems/clone-graph/description/

/**
 * Definition for Node.
 * class Node {
 *     val: number
 *     neighbors: Node[]
 *     constructor(val?: number, neighbors?: Node[]) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.neighbors = (neighbors===undefined ? [] : neighbors)
 *     }
 * }
 */

function _cloneGraph(node, lookup): Node | null {
  if (lookup[node.val.toString()]) return lookup[node.val];
  const clone =  new Node(node.val, []);
  lookup[`${clone.val.toString()}`] = clone;
  clone.neighbors = node.neighbors.map(n => _cloneGraph(n, lookup));
  return clone;
}

function cloneGraph(node: Node | null): Node | null { 
  if (!node) return null;
  if (node.val === 1 && node.neighbors.length === 0) return new Node(1, []);  
  const dict = {};
  const clone = _cloneGraph(node, dict);  
  return clone
};

