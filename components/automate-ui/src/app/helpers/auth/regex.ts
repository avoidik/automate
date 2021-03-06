export class Regex {

  public static readonly patterns = {

    // all auth IDs use this pattern
    // Note: the restriction to <= 64 chars is implemented by another validator,
    // so that we get a nicer error message. Here it's enough to restrict the
    // char set.
    // Interestingly, the regex validator doesn't seem to fire for empty strings.
    ID: '[0-9a-z-_]+',

    // NB: neither \S nor ^\s work inside the brackets in this regex language.
    NON_BLANK: '.*[^ ].*',

    // Only allows wildcard alone or words and numbers, but not combined
    // Legal Values: *, chef, _state, etc.
    // Illegal Values: abc*, *chef, c*h*e*f, **

    // Allows no special characters except hyphen
    NO_MIXED_WILDCARD_ALLOW_HYPHEN: '^(\\*|[-\\w]+)$',

    // Allows all special characters except colon :
    NO_MIXED_WILDCARD_ALLOW_SPECIAL: '^(\\*|[^:*]+)$'
  };

}
