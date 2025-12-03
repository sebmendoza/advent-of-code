import sys
import importlib
import importlib.util
from pathlib import Path


def main():
    days = sys.argv[1:]
    for d in days:
        p = Path(__file__).parent / d / f"{d}.py"
        spec = importlib.util.spec_from_file_location(d, str(p))
        mod = importlib.util.module_from_spec(spec)
        spec.loader.exec_module(mod)
        daymodule = getattr(mod, "run", None) or getattr(
            mod, "main", None) or getattr(mod, d, None)
        daymodule()


if __name__ == "__main__":
    main()
