## Subdomain: Products

Bounded Context: ViewProductsList
    Input: Category
    Output: ProductsList or EmptyProductsList
    Dependencies: CheckCategoryExist, CheckProductsExist

    Steps:
    1. Click or search a category item
    2. Validate category id if exist
    3. Get all products associated with the category id
    4. Return and load the products