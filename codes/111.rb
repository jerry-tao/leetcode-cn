def min_depth(root)
    return 0 if root.nil?
    return 1 if root.left.nil? && root.right.nil?
    return min_depth(root.left)+1 if root.right.nil?
    return min_depth(root.right)+1 if root.left.nil?
    return [min_depth(root.left)+1, min_depth(root.right)+1].min
end
