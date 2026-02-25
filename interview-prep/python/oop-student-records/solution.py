"""
OOP Challenge: Student Records

Part 1 — Student dataclass
    Create a Student class using @dataclass with fields: name (str), student_id (int),
    grades (list[float]). Implement:
    - gpa property: average of grades, 0.0 if no grades
    - __repr__ and __str__: repr for debugging, str for display
    - __eq__: two students are equal if same student_id
    - __hash__: based on student_id so students can go in sets/dicts
    - __lt__: default comparison by student_id (needed for sorted())

Part 2 — Nationality subclasses
    Create subclasses: ChineseStudent, AmericanStudent, IndianStudent.
    Each overrides a `country` property returning the country name.
    Implement a factory function `create_student(name, student_id, grades)` that
    returns the correct subclass based on last name heuristics:
    - Last name in ("Li", "Wang", "Zhang", "Chen", "Liu") → ChineseStudent
    - Last name in ("Smith", "Johnson", "Williams", "Brown", "Jones") → AmericanStudent
    - Last name in ("Singh", "Kumar", "Patel", "Sharma", "Gupta") → IndianStudent
    - Otherwise → base Student

Part 3 — Multi-criteria sorting
    Implement these sorting helpers (can be standalone functions or classmethods):
    - sort_by_name(students): alphabetical by name
    - sort_by_gpa(students): descending by GPA, ties broken by name ascending
    - sort_by_id(students): ascending by student_id

Pythonic goals:
- @dataclass with field() for mutable default (grades)
- @property for computed attributes (gpa, country)
- __eq__ / __hash__ / __lt__ dunder methods
- Inheritance with super().__init__()
- sorted() with key=lambda for multi-criteria sorting
"""

from dataclasses import dataclass, field
from typing import List


# Part 1: Student dataclass
@dataclass(order=True, frozen=False)
class Student:
    # standard field
    name: str
    student_id: int
    grades: list[float] = field(default_factory=list)

    # property for GPA
    @property
    def gpa(self) -> float:
        return sum(self.grades) / len(self.grades) if self.grades else 0.0

    def __hash__(self):
        return hash(self.student_id)


# Part 2: Nationality subclasses + factory
# TODO: implement


# Part 3: Sorting helpers
# TODO: implement


if __name__ == "__main__":
    # ── Part 1: Basic Student ──

    s1 = Student("Alice Smith", 101, [3.5, 3.8, 4.0])
    s2 = Student("Bob Johnson", 102, [2.9, 3.2, 3.0])
    s3 = Student("Alice Smith", 101, [3.5, 3.8, 4.0])
    s4 = Student("Charlie Wang", 103, [])

    # GPA
    assert abs(s1.gpa - 3.7667) < 0.001, f"s1 gpa: {s1.gpa}"
    assert abs(s2.gpa - 3.0333) < 0.001, f"s2 gpa: {s2.gpa}"
    assert s4.gpa == 0.0, "empty grades should give 0.0 GPA"

    # __eq__ based on student_id, not name/grades
    assert s1 == s3, "same student_id should be equal"
    assert s1 != s2, "different student_id should not be equal"

    # __hash__ — students can go in sets
    student_set = {s1, s2, s3}
    assert len(student_set) == 2, "s1 and s3 have same id, set should deduplicate"

    # __repr__ should include class name and key fields
    r = repr(s1)
    assert "Alice Smith" in r and "101" in r, f"repr missing info: {r}"

    # __str__ should be human-readable
    st = str(s1)
    assert "Alice Smith" in st, f"str missing name: {st}"

    # __lt__ by student_id
    assert s1 < s2, "101 < 102"
    assert not s2 < s1
    assert sorted([s2, s1]) == [s1, s2], "sorted by student_id"

    print("Part 1 passed!")

    # ── Part 2: Nationality subclasses ──

    c1 = create_student("Wei Li", 201, [3.9, 4.0])
    c2 = create_student("John Smith", 202, [3.5, 3.6])
    c3 = create_student("Raj Singh", 203, [3.7, 3.8])
    c4 = create_student("Unknown Person", 204, [3.0])

    assert isinstance(c1, ChineseStudent), f"expected ChineseStudent, got {type(c1)}"
    assert isinstance(c2, AmericanStudent), f"expected AmericanStudent, got {type(c2)}"
    assert isinstance(c3, IndianStudent), f"expected IndianStudent, got {type(c3)}"
    assert isinstance(c4, Student), f"expected base Student, got {type(c4)}"
    # subclasses should also be Student instances
    assert isinstance(c1, Student), "ChineseStudent should be a Student"

    assert c1.country == "China"
    assert c2.country == "America"
    assert c3.country == "India"
    assert c4.country == "Unknown"

    # Subclass students should still support all Part 1 features
    assert abs(c1.gpa - 3.95) < 0.001
    assert c1 != c2  # different IDs

    print("Part 2 passed!")

    # ── Part 3: Multi-criteria sorting ──

    students = [
        Student("Charlie Brown", 105, [3.5, 3.5]),       # GPA 3.5
        Student("Alice Zhang", 101, [4.0, 3.8]),          # GPA 3.9
        Student("Bob Kumar", 103, [3.9, 3.9]),            # GPA 3.9
        Student("Diana Smith", 102, [2.5, 3.0]),          # GPA 2.75
        Student("Eve Li", 104, [3.5, 3.5]),               # GPA 3.5
    ]

    # sort_by_name: alphabetical
    by_name = sort_by_name(students)
    names = [s.name for s in by_name]
    assert names == ["Alice Zhang", "Bob Kumar", "Charlie Brown", "Diana Smith", "Eve Li"], \
        f"by name: {names}"

    # sort_by_gpa: descending GPA, ties broken by name ascending
    by_gpa = sort_by_gpa(students)
    gpa_names = [(s.name, round(s.gpa, 2)) for s in by_gpa]
    assert gpa_names == [
        ("Alice Zhang", 3.9),
        ("Bob Kumar", 3.9),
        ("Charlie Brown", 3.5),
        ("Eve Li", 3.5),
        ("Diana Smith", 2.75),
    ], f"by gpa: {gpa_names}"

    # sort_by_id: ascending
    by_id = sort_by_id(students)
    ids = [s.student_id for s in by_id]
    assert ids == [101, 102, 103, 104, 105], f"by id: {ids}"

    # Original list should NOT be modified
    assert students[0].name == "Charlie Brown", "sorting should not mutate original"

    print("Part 3 passed!")
    print("All tests passed!")
