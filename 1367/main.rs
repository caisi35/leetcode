use std::rc::Rc;
use std::cell::RefCell;

struct Solution {

}

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}
// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
  pub val: i32,
  pub left: Option<Rc<RefCell<TreeNode>>>,
  pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
  #[inline]
  pub fn new(val: i32) -> Self {
    TreeNode {
      val,
      left: None,
      right: None
    }
  }
}

impl Solution {
    pub fn is_sub_path(head: Option<Box<ListNode>>, root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        match (head.clone(), root.clone()) {
            (None, _) => true,
            (_, None) => false,
            (Some(list_node), Some(tree_node)) => {
                Self::dfs(&head, &root) ||
                Self::is_sub_path(head.clone(), tree_node.borrow().right.clone()) ||
                Self::is_sub_path(head.clone(), tree_node.borrow().left.clone())
            }
        }
    }

    pub fn dfs(head: &Option<Box<ListNode>>, root: &Option<Rc<RefCell<TreeNode>>>) -> bool {
        match (head, root) {
            (None, _) => true,
            (_, None) => false,
            (Some(list_node), Some(tree_node)) => {
                list_node.val == tree_node.borrow().val &&
                (Self::dfs(&list_node.next, &tree_node.borrow().left) ||
                (Self::dfs(&list_node.next, &tree_node.borrow().right)))
            }
        }
    }
}

fn main() {
    let head = Box::new(ListNode::new(0));
    let root = Rc::new(RefCell::new(TreeNode::new(0)));
    println!("{}", Solution::is_sub_path(Some(head), Some(root)));
}